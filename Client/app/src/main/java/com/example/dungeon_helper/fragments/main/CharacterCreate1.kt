package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentCharacterCreate1Binding

class CharacterCreate1 : Fragment() {

    companion object {
        fun newInstance() = CharacterCreate1()
    }

    private lateinit var viewModel: CharacterCreate1ViewModel
    private var _binding: FragmentCharacterCreate1Binding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val characterCreate1ViewModel = ViewModelProvider(this)[CharacterCreate1ViewModel::class.java]
        _binding = FragmentCharacterCreate1Binding.inflate(inflater,container,false)
        val root: View = binding.root
        val textView: TextView = binding.textCharacterCreate1
        characterCreate1ViewModel.text.observe(viewLifecycleOwner){
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
            (requireActivity() as MainActivity).showConfirmationDialog(
                "Подтверждение возврата",
                "Данные не сохранены. Вы уверены, что хотите вернуться?",
                {
                    val activity = requireActivity() as MainActivity
                    val navView = activity.getNavView()
                    val menu = navView.menu
                    val menuItem1 = menu.findItem(R.id.navigation_info)
                    val menuItem2 = menu.findItem(R.id.navigation_character)
                    val menuItem3 = menu.findItem(R.id.navigation_lobby)
                    val menuItem4 = menu.findItem(R.id.navigation_account)
                    menuItem1.isVisible = true
                    menuItem2.isVisible = true
                    menuItem3.isVisible = true
                    menuItem4.isVisible = true
                    activity.navController.navigate(R.id.action_characterCreate1_to_navigation_character)
                },
                {  }
            )
        }
        val forwBtn = binding.forwBtn
        forwBtn.setOnClickListener{
            (activity as MainActivity).navController.navigate(R.id.action_characterCreate1_to_characterCreate2)
        }
    }

}