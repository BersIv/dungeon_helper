package com.example.dungeon_helper.fragments.main.character

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.AdapterView
import android.widget.ArrayAdapter
import android.widget.TextView
import android.widget.Toast
import androidx.lifecycle.Observer
import com.google.android.material.button.MaterialButton
import com.google.android.material.textview.MaterialTextView
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentCharacterCreate4Binding
import com.example.dungeon_helper.shared.JsonHelper
import com.example.dungeon_helper.shared.SharedViewModel
import com.example.dungeon_helper.shared.Stats
import com.example.dungeon_helper.shared.Subrace
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.RequestBody
import okhttp3.RequestBody.Companion.toRequestBody
import org.json.JSONObject
import java.io.IOException

class CharacterCreate4 : Fragment() {

    companion object {
        fun newInstance() = CharacterCreate4()
    }

    private lateinit var viewModel: CharacterCreate4ViewModel
    private var _binding: FragmentCharacterCreate4Binding? = null
    private val binding get() = _binding!!

    private lateinit var currentCharacteristicButton: MaterialButton
    private lateinit var currentSumTextView: TextView
    private lateinit var currentModTextView: TextView
    private lateinit var currentSpentPointsTextView: TextView
    private lateinit var valueTextView: MaterialTextView
    private var characteristicSelected: Boolean = false

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
       val characterCreate4ViewModel = ViewModelProvider(this)[CharacterCreate4ViewModel::class.java]
        _binding = FragmentCharacterCreate4Binding.inflate(inflater, container, false)
        val root: View = binding.root
        val textView: TextView = binding.textCharacterCreate4
        characterCreate4ViewModel.text.observe(viewLifecycleOwner){
            textView.text = it
        }
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    private fun updateCharacteristic(newText: String, valueTextView: MaterialTextView) {
        binding.Characteristic.text = newText
        valueTextView.text = currentCharacteristicButton.text
    }

    private fun enablePlusMinusButtons() {
        // Разблокировать кнопки минус и плюс только если выбрана характеристика
        binding.plusBtn.isEnabled = characteristicSelected
        binding.minusBtn.isEnabled = characteristicSelected
    }
    private fun setCurrentCharacteristicButton(
        button: MaterialButton,
        sumTextView: TextView,
        modTextView: TextView,
        spentPointsTextView: TextView,
        valueTextView: MaterialTextView,
        textTextView: TextView
    ) {
        currentCharacteristicButton = button
        currentSumTextView = sumTextView
        currentModTextView = modTextView
        currentSpentPointsTextView = spentPointsTextView
        this.valueTextView = valueTextView

        updateCharacteristic(textTextView.text.toString(), valueTextView)
        enablePlusMinusButtons()
    }

    private var currentPoints = 0
    private fun handlePlusButtonClick(
        button: MaterialButton,
        valueTextView: MaterialTextView,
        sumTextView: TextView,
        modTextView: TextView,
        spentPointsTextView: TextView,
        textTextView: TextView
    ) {
        val currentValue = button.text.toString().toInt()

        val newValue = currentValue + 1
        valueTextView.text = newValue.toString()

        button.text = newValue.toString()
        sumTextView.text = newValue.toString()
        val modifierValue = (newValue - 10) / 2
        modTextView.text = modifierValue.toString()
        currentPoints += 1
        spentPointsTextView.text = currentPoints.toString()
        updateCharacteristic(textTextView.text.toString(), valueTextView)
    }

    private fun handleMinusButtonClick(
        button: MaterialButton,
        valueTextView: MaterialTextView,
        sumTextView: TextView,
        modTextView: TextView,
        spentPointsTextView: TextView,
        textTextView: TextView
    ) {
        val currentValue = button.text.toString().toInt()

        val newValue = currentValue - 1

        valueTextView.text = newValue.toString()

        button.text = newValue.toString()

        sumTextView.text = newValue.toString()

        val modifierValue = (newValue - 10) / 2
        modTextView.text = modifierValue.toString()

        currentPoints -= 1
        spentPointsTextView.text = currentPoints.toString()

        updateCharacteristic(textTextView.text.toString(), valueTextView)
    }

    override fun onStart() {
        super.onStart()

        val sharedViewModel = ViewModelProvider(requireActivity())[SharedViewModel::class.java]

        val langCharacter = binding.LangCharacter.editText
        val optionsSpinner = binding.spinnerOptions
        val backBtn = binding.backBtn
        val forwBtn = binding.forwBtn
        var selected = ""

        var listSubraces = sharedViewModel.subraces.value
        sharedViewModel.subraces.observe(viewLifecycleOwner, Observer {
            // updating data in displayMsg
            listSubraces = it
        })

        val arrayOptions: ArrayList<String> = ArrayList()

        if (listSubraces != null) {
            for (i in listSubraces!!) {
                arrayOptions.add(i.raceName)
            }
        }


        val optionsArray = arrayOf("Высший", "Лесной", "Темный")

        val adapter =
            ArrayAdapter(requireContext(), android.R.layout.simple_spinner_item, arrayOptions)

        adapter.setDropDownViewResource(android.R.layout.simple_spinner_dropdown_item)


        optionsSpinner.adapter = adapter


        optionsSpinner.onItemSelectedListener = object : AdapterView.OnItemSelectedListener {
            override fun onItemSelected(
                parentView: AdapterView<*>?,
                selectedItemView: View?,
                position: Int,
                id: Long
            ) {

                val selectedItem = parentView?.getItemAtPosition(position).toString()
                Toast.makeText(requireContext(), "Выбрано: $selectedItem", Toast.LENGTH_SHORT)
                    .show()

                selected = selectedItem
            }

            override fun onNothingSelected(parentView: AdapterView<*>?) {
                forwBtn.isEnabled = false

            }

        }

        val valueTextView = binding.value
        val pointsText = binding.points

        val silBtn = binding.sil
        val textSilText = binding.textSil
        val sumSil = binding.sumSil
        val modSil = binding.modSil

        silBtn.setOnClickListener {
            characteristicSelected = true
            setCurrentCharacteristicButton(silBtn, sumSil, modSil, pointsText, valueTextView, textSilText)
        }
        val lovBtn = binding.lov
        val textLovText = binding.textLov
        val sumLov = binding.sumLov
        val modLov = binding.modLov

        lovBtn.setOnClickListener {
            characteristicSelected = true
            setCurrentCharacteristicButton(lovBtn, sumLov, modLov, pointsText, valueTextView, textLovText)
        }

        val telBtn = binding.tel
        val textTelText = binding.textTel
        val sumTel = binding.sumTel
        val modTel = binding.modTel
        telBtn.setOnClickListener {
            characteristicSelected = true
            setCurrentCharacteristicButton(telBtn,sumTel, modTel, pointsText, valueTextView,textTelText)
        }
        val intBtn = binding.intel
        val textIntText = binding.textIntel
        val sumInt = binding.sumIntel
        val modInt = binding.modIntel
        intBtn.setOnClickListener {
            characteristicSelected = true
            setCurrentCharacteristicButton(intBtn, sumInt, modInt, pointsText, valueTextView, textIntText)
        }
        val mdrBtn = binding.mdr
        val textMdrText = binding.textMdr
        val sumMdr = binding.sumMdr
        val modMdr = binding.modMdr
        mdrBtn.setOnClickListener {
            characteristicSelected = true
            setCurrentCharacteristicButton(mdrBtn, sumMdr, modMdr, pointsText, valueTextView, textMdrText)
        }
        val harBtn = binding.har
        val textHarText = binding.textHar
        val sumHar = binding.sumHar
        val modHar = binding.modHar
        harBtn.setOnClickListener {
            characteristicSelected = true
            setCurrentCharacteristicButton(harBtn, sumHar, modHar,pointsText,valueTextView,textHarText)
        }

        val plusBtn = binding.plusBtn
        plusBtn.isEnabled = false
        plusBtn.setOnClickListener {
               handlePlusButtonClick(
                   currentCharacteristicButton,
                   valueTextView,
                   currentSumTextView,
                   currentModTextView,
                   currentSpentPointsTextView,
                   textSilText
               )

        }

        val minusBtn = binding.minusBtn
        minusBtn.isEnabled = false
        minusBtn.setOnClickListener {
            handleMinusButtonClick(
                currentCharacteristicButton,
                valueTextView,
                currentSumTextView,
                currentModTextView,
                currentSpentPointsTextView,
                textSilText
            )
        }





        backBtn.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_characterCreate4_to_characterCreate3)
        }

        forwBtn.setOnClickListener {
            GlobalScope.launch(Dispatchers.Main) {

                var subrace = Subrace(0, "", Stats(0, 0, 0, 0, 0, 0))

                if (listSubraces != null) {
                    for (i in listSubraces!!) {
                        if (selected == i.raceName) {
                            subrace = i
                        }
                    }
                }

                val jsonBody = JSONObject().apply {
                    put("id", subrace.id)
                }

                println(jsonBody)

                val mediaType = "application/json; charset=utf-8".toMediaType()
                val body: RequestBody = jsonBody.toString().toRequestBody(mediaType)

                println(body)

                var token: String = ""
                sharedViewModel.token.observe(viewLifecycleOwner, Observer {
                    // updating data in displayMsg
                    token = it
                })


                //println(getCookieValue())

                val client = OkHttpClient.Builder().addInterceptor { chain ->
                    val original = chain.request()
                    val authorized = original.newBuilder()
                        .header("Authorization", token)
                        .build()
                    chain.proceed(authorized)
                }
                    .build()

                val request1 = Request.Builder()
                    .url("http://194.247.187.44:5000/skills/getSkills")
                    .get()
                    .build()

                val request = Request.Builder()
                    .url("http://194.247.187.44:5000/alignments/getAlignments")
                    .get()
                    .build()

                try {
                    val response1 = withContext(Dispatchers.IO) {
                        client.newCall(request1).execute()
                    }
                    val response = withContext(Dispatchers.IO) {
                        client.newCall(request).execute()
                    }

                    if (!response1.isSuccessful) {
                        throw IOException("Запрос к серверу не был успешен:" +
                                " ${response1.code} ${response1.message}")
                    }

                    if (!response.isSuccessful) {
                        throw IOException("Запрос к серверу не был успешен:" +
                                " ${response.code} ${response.message}")
                    }

                    println("${response1.code} ${response1.message}")

                    val deserializedASkillList = JsonHelper.deserializeListFromJsonSkillList(response1.body!!.string())
                    println("Десериализованный объект:\n$deserializedASkillList")

                    sharedViewModel.skills.value = deserializedASkillList

                    println("${response.code} ${response.message}")

                    val deserializedAlignmentList = JsonHelper.deserializeListFromJsonAlignmentList(response.body!!.string())
                    println("Десериализованный объект:\n$deserializedAlignmentList")

                    sharedViewModel.alignments.value = deserializedAlignmentList

                    var subrace = Subrace(0, "", Stats(0, 0, 0, 0, 0, 0))

                    if (listSubraces != null) {
                        for (i in listSubraces!!) {
                            if (selected == i.raceName) {
                                subrace = i
                            }
                        }
                    }

                    val newCharacter = sharedViewModel.newCharacter.value

                    newCharacter?.subrace = subrace
                    newCharacter?.stats?.strength = sumSil.text.toString().toInt()
                    newCharacter?.stats?.dexterity = sumLov.text.toString().toInt()
                    newCharacter?.stats?.constitution = sumTel.text.toString().toInt()
                    newCharacter?.stats?.intelligence = sumInt.text.toString().toInt()
                    newCharacter?.stats?.wisdom = sumMdr.text.toString().toInt()
                    newCharacter?.stats?.charisma = sumHar.text.toString().toInt()
                    newCharacter?.addLanguage = langCharacter?.text.toString()

                    (activity as MainActivity).navController.navigate(R.id.action_characterCreate4_to_characterCreate5)

                } catch (e: IOException) {
                    println("Ошибка подключения: $e");
                }
            }


        }



    }


}