package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import com.example.dungeon_helper.R
import android.widget.TextView
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.databinding.FragmentCharacterMainBinding

class CharacterMain : Fragment() {

    companion object {
        fun newInstance() = CharacterMain()
    }
    private lateinit var viewModel: CharacterMainViewModel


    private var _binding: FragmentCharacterMainBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View {
        val characterMainViewModel = ViewModelProvider(this)[CharacterMainViewModel::class.java]

        _binding = FragmentCharacterMainBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textCharacter
        characterMainViewModel.text.observe(viewLifecycleOwner) {
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
        val createCharBtn = binding.createCharBtn
        createCharBtn.setOnClickListener{
            val activity = requireActivity() as MainActivity
            val navView = activity.getNavView()
            val menu = navView.menu
            val menuItem1 = menu.findItem(R.id.navigation_info)
            val menuItem2 = menu.findItem(R.id.navigation_character)
            val menuItem3 = menu.findItem(R.id.navigation_lobby)
            val menuItem4 = menu.findItem(R.id.navigation_account)
            menuItem1.isVisible = false
            menuItem2.isVisible = false
            menuItem3.isVisible = false
            menuItem4.isVisible = false
            (activity as MainActivity).navController.navigate(R.id.action_navigation_character_to_characterCreate1)
        }

    }

}