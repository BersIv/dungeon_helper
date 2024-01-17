package com.example.dungeon_helper.fragments.main

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
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentCharacterCreate5Binding

class CharacterCreate5 : Fragment() {

    companion object {
        fun newInstance() = CharacterCreate5()
    }

    private lateinit var viewModel: CharacterCreate5ViewModel
    private var _binding: FragmentCharacterCreate5Binding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val characterCreate5ViewModel = ViewModelProvider(this)[CharacterCreate5ViewModel::class.java]
        _binding = FragmentCharacterCreate5Binding.inflate(inflater, container, false)
        val root: View = binding.root
        val textView: TextView = binding.textCharacterCreate5
        characterCreate5ViewModel.text.observe(viewLifecycleOwner){
            textView.text = it
        }
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()
        val backBtn = binding.backBtn
        backBtn.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_characterCreate5_to_characterCreate4)
        }
        val forwBtn = binding.forwBtn
        forwBtn.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_characterCreate5_to_characterCreate6)
        }

        val optionsSpinner = binding.spinnerOptions
        val optionsArray = arrayOf("Законно-добрый", "Нейтрально-добрый", "Хаотично-добрый","Законно-нейтральные",
            "Нейтральный","Хаотично-нейтральный", "Законно-злой", "Нейтрально-злой","Хаотично-злой")

        val adapter =
            ArrayAdapter(requireContext(), android.R.layout.simple_spinner_item, optionsArray)

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
            }

            override fun onNothingSelected(parentView: AdapterView<*>?) {
                forwBtn.isEnabled = false

            }

        }

    }

}